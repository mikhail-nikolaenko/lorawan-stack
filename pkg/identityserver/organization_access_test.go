// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package identityserver

import (
	"testing"

	pbtypes "github.com/gogo/protobuf/types"
	"github.com/smartystreets/assertions/should"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/storetest"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"google.golang.org/grpc"
)

func TestOrganizationAPIKeys(t *testing.T) {
	p := &storetest.Population{}

	admin := p.NewUser()
	admin.Admin = true
	adminKey, _ := p.NewAPIKey(admin.GetEntityIdentifiers(), ttnpb.Right_RIGHT_ALL)
	adminCreds := rpcCreds(adminKey)

	usr1 := p.NewUser()
	usr1Key, _ := p.NewAPIKey(usr1.GetEntityIdentifiers(), ttnpb.Right_RIGHT_ALL)
	usr1Creds := rpcCreds(usr1Key)

	org1 := p.NewOrganization(usr1.GetOrganizationOrUserIdentifiers())
	limitedKey, _ := p.NewAPIKey(usr1.GetEntityIdentifiers(),
		ttnpb.Right_RIGHT_ORGANIZATION_INFO,
		ttnpb.Right_RIGHT_ORGANIZATION_SETTINGS_BASIC,
		ttnpb.Right_RIGHT_ORGANIZATION_SETTINGS_API_KEYS,
	)
	limitedCreds := rpcCreds(limitedKey)

	orgKey, _ := p.NewAPIKey(org1.GetEntityIdentifiers(),
		ttnpb.Right_RIGHT_ORGANIZATION_INFO,
		ttnpb.Right_RIGHT_ORGANIZATION_APPLICATIONS_LIST,
	)
	orgCreds := rpcCreds(orgKey)

	t.Parallel()
	a, ctx := test.New(t)

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		is.config.AdminRights.All = true

		reg := ttnpb.NewOrganizationAccessClient(cc)

		// GetAPIKey that doesn't exist.
		got, err := reg.GetAPIKey(ctx, &ttnpb.GetOrganizationAPIKeyRequest{
			OrganizationIds: org1.GetIds(),
			KeyId:           "does-not-exist",
		}, limitedCreds)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsNotFound(err), should.BeTrue)
		}
		a.So(got, should.BeNil)

		// UpdateAPIKey that doesn't exist.
		updated, err := reg.UpdateAPIKey(ctx, &ttnpb.UpdateOrganizationAPIKeyRequest{
			OrganizationIds: org1.GetIds(),
			ApiKey: &ttnpb.APIKey{
				Id: "does-not-exist",
			},
			FieldMask: &pbtypes.FieldMask{Paths: []string{"name"}},
		}, limitedCreds)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsNotFound(err), should.BeTrue)
		}
		a.So(updated, should.BeNil)

		// CreateAPIKey with rights that caller doesn't have.
		apiKey, err := reg.CreateAPIKey(ctx, &ttnpb.CreateOrganizationAPIKeyRequest{
			OrganizationIds: org1.GetIds(),
			Name:            "api-key-name",
			Rights:          []ttnpb.Right{ttnpb.Right_RIGHT_ORGANIZATION_ALL},
		}, limitedCreds)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}
		a.So(apiKey, should.BeNil)

		// UpdateAPIKey adding rights that caller doesn't have.
		updated, err = reg.UpdateAPIKey(ctx, &ttnpb.UpdateOrganizationAPIKeyRequest{
			OrganizationIds: org1.GetIds(),
			ApiKey: &ttnpb.APIKey{
				Id: orgKey.GetId(),
				Rights: []ttnpb.Right{
					ttnpb.Right_RIGHT_ORGANIZATION_INFO,
					ttnpb.Right_RIGHT_ORGANIZATION_APPLICATIONS_LIST,
					ttnpb.Right_RIGHT_ORGANIZATION_DELETE,
				},
			},
			FieldMask: &pbtypes.FieldMask{Paths: []string{"rights"}},
		}, limitedCreds)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}
		a.So(updated, should.BeNil)

		// UpdateAPIKey removing rights that caller doesn't have.
		updated, err = reg.UpdateAPIKey(ctx, &ttnpb.UpdateOrganizationAPIKeyRequest{
			OrganizationIds: org1.GetIds(),
			ApiKey: &ttnpb.APIKey{
				Id: orgKey.GetId(),
				Rights: []ttnpb.Right{
					ttnpb.Right_RIGHT_ORGANIZATION_INFO,
				},
			},
			FieldMask: &pbtypes.FieldMask{Paths: []string{"rights"}},
		}, limitedCreds)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}
		a.So(updated, should.BeNil)

		// UpdateAPIKey removing rights that caller has and adding rights that caller has.
		updated, err = reg.UpdateAPIKey(ctx, &ttnpb.UpdateOrganizationAPIKeyRequest{
			OrganizationIds: org1.GetIds(),
			ApiKey: &ttnpb.APIKey{
				Id: orgKey.GetId(),
				Rights: []ttnpb.Right{
					ttnpb.Right_RIGHT_ORGANIZATION_SETTINGS_BASIC,
					ttnpb.Right_RIGHT_ORGANIZATION_APPLICATIONS_LIST,
				},
			},
			FieldMask: &pbtypes.FieldMask{Paths: []string{"rights"}},
		}, limitedCreds)
		if a.So(err, should.BeNil) && a.So(updated, should.NotBeNil) {
			a.So(updated.Rights, should.Resemble, []ttnpb.Right{
				ttnpb.Right_RIGHT_ORGANIZATION_SETTINGS_BASIC,
				ttnpb.Right_RIGHT_ORGANIZATION_APPLICATIONS_LIST,
			})
		}

		// API Key CRUD with different invalid credentials.
		for _, opts := range [][]grpc.CallOption{nil, {orgCreds}} {
			created, err := reg.CreateAPIKey(ctx, &ttnpb.CreateOrganizationAPIKeyRequest{
				OrganizationIds: org1.GetIds(),
				Name:            "api-key-name",
				Rights:          []ttnpb.Right{ttnpb.Right_RIGHT_ORGANIZATION_INFO},
			}, opts...)
			if a.So(err, should.NotBeNil) && a.So(errors.IsPermissionDenied(err), should.BeTrue) {
				a.So(created, should.BeNil)
			}

			list, err := reg.ListAPIKeys(ctx, &ttnpb.ListOrganizationAPIKeysRequest{
				OrganizationIds: org1.GetIds(),
			}, opts...)
			if a.So(err, should.NotBeNil) && a.So(errors.IsPermissionDenied(err), should.BeTrue) {
				a.So(list, should.BeNil)
			}

			got, err := reg.GetAPIKey(ctx, &ttnpb.GetOrganizationAPIKeyRequest{
				OrganizationIds: org1.GetIds(),
				KeyId:           orgKey.GetId(),
			}, opts...)
			if a.So(err, should.NotBeNil) && a.So(errors.IsPermissionDenied(err), should.BeTrue) {
				a.So(got, should.BeNil)
			}

			updated, err := reg.UpdateAPIKey(ctx, &ttnpb.UpdateOrganizationAPIKeyRequest{
				OrganizationIds: org1.GetIds(),
				ApiKey: &ttnpb.APIKey{
					Id:   orgKey.GetId(),
					Name: "api-key-name-updated",
				},
				FieldMask: &pbtypes.FieldMask{Paths: []string{"name"}},
			}, opts...)
			if a.So(err, should.NotBeNil) && a.So(errors.IsPermissionDenied(err), should.BeTrue) {
				a.So(updated, should.BeNil)
			}
		}

		// API Key CRUD with different valid credentials.
		for _, opts := range [][]grpc.CallOption{{adminCreds}, {usr1Creds}, {limitedCreds}} {
			created, err := reg.CreateAPIKey(ctx, &ttnpb.CreateOrganizationAPIKeyRequest{
				OrganizationIds: org1.GetIds(),
				Name:            "api-key-name",
				Rights:          []ttnpb.Right{ttnpb.Right_RIGHT_ORGANIZATION_INFO},
			}, opts...)
			if a.So(err, should.BeNil) && a.So(created, should.NotBeNil) {
				a.So(created.Name, should.Equal, "api-key-name")
				a.So(created.Rights, should.Resemble, []ttnpb.Right{ttnpb.Right_RIGHT_ORGANIZATION_INFO})
			}

			list, err := reg.ListAPIKeys(ctx, &ttnpb.ListOrganizationAPIKeysRequest{
				OrganizationIds: org1.GetIds(),
			}, opts...)
			if a.So(err, should.BeNil) && a.So(list, should.NotBeNil) && a.So(list.ApiKeys, should.HaveLength, 2) {
				for _, k := range list.ApiKeys {
					if k.Id == created.Id {
						a.So(k.Name, should.Resemble, created.Name)
					}
				}
			}

			got, err := reg.GetAPIKey(ctx, &ttnpb.GetOrganizationAPIKeyRequest{
				OrganizationIds: org1.GetIds(),
				KeyId:           created.GetId(),
			}, opts...)
			if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) {
				a.So(got.Name, should.Equal, created.Name)
			}

			updated, err := reg.UpdateAPIKey(ctx, &ttnpb.UpdateOrganizationAPIKeyRequest{
				OrganizationIds: org1.GetIds(),
				ApiKey: &ttnpb.APIKey{
					Id:   created.GetId(),
					Name: "api-key-name-updated",
				},
				FieldMask: &pbtypes.FieldMask{Paths: []string{"name"}},
			}, opts...)
			if a.So(err, should.BeNil) && a.So(updated, should.NotBeNil) {
				a.So(updated.Name, should.Equal, "api-key-name-updated")
			}

			deleted, err := reg.UpdateAPIKey(ctx, &ttnpb.UpdateOrganizationAPIKeyRequest{
				OrganizationIds: org1.GetIds(),
				ApiKey: &ttnpb.APIKey{
					Id: created.GetId(),
				},
				FieldMask: &pbtypes.FieldMask{Paths: []string{"rights"}},
			}, opts...)
			if a.So(err, should.BeNil) && a.So(deleted, should.NotBeNil) {
				a.So(deleted.Rights, should.BeNil)
			}

			got, err = reg.GetAPIKey(ctx, &ttnpb.GetOrganizationAPIKeyRequest{
				OrganizationIds: org1.GetIds(),
				KeyId:           created.GetId(),
			}, opts...)
			if a.So(err, should.NotBeNil) {
				a.So(errors.IsNotFound(err), should.BeTrue)
			}
			a.So(got, should.BeNil)
		}
	}, withPrivateTestDatabase(p))
}

func TestOrganizationCollaborators(t *testing.T) {
	p := &storetest.Population{}

	admin := p.NewUser()
	admin.Admin = true
	adminKey, _ := p.NewAPIKey(admin.GetEntityIdentifiers(), ttnpb.Right_RIGHT_ALL)
	adminCreds := rpcCreds(adminKey)

	usr1 := p.NewUser()
	usr1Key, _ := p.NewAPIKey(usr1.GetEntityIdentifiers(), ttnpb.Right_RIGHT_ALL)
	usr1Creds := rpcCreds(usr1Key)

	org1 := p.NewOrganization(usr1.GetOrganizationOrUserIdentifiers())

	limitedKey, _ := p.NewAPIKey(usr1.GetEntityIdentifiers(),
		ttnpb.Right_RIGHT_ORGANIZATION_INFO,
		ttnpb.Right_RIGHT_ORGANIZATION_SETTINGS_BASIC,
		ttnpb.Right_RIGHT_ORGANIZATION_SETTINGS_MEMBERS,
	)
	limitedCreds := rpcCreds(limitedKey)

	orgKey, _ := p.NewAPIKey(org1.GetEntityIdentifiers(),
		ttnpb.Right_RIGHT_ORGANIZATION_INFO,
		ttnpb.Right_RIGHT_ORGANIZATION_APPLICATIONS_LIST,
	)
	orgCreds := rpcCreds(orgKey)

	usr2 := p.NewUser()
	p.NewMembership(
		usr2.GetOrganizationOrUserIdentifiers(),
		org1.GetEntityIdentifiers(),
		ttnpb.Right_RIGHT_ORGANIZATION_INFO,
		ttnpb.Right_RIGHT_ORGANIZATION_APPLICATIONS_LIST,
	)

	usr3 := p.NewUser()

	t.Parallel()
	a, ctx := test.New(t)

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		is.config.AdminRights.All = true

		reg := ttnpb.NewOrganizationAccessClient(cc)

		// GetCollaborator that doesn't exist.
		got, err := reg.GetCollaborator(ctx, &ttnpb.GetOrganizationCollaboratorRequest{
			OrganizationIds: org1.GetIds(),
			Collaborator:    usr3.GetOrganizationOrUserIdentifiers(),
		}, adminCreds)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsNotFound(err), should.BeTrue)
		}
		a.So(got, should.BeNil)

		// SetCollaborator adding rights that caller doesn't have.
		_, err = reg.SetCollaborator(ctx, &ttnpb.SetOrganizationCollaboratorRequest{
			OrganizationIds: org1.GetIds(),
			Collaborator: &ttnpb.Collaborator{
				Ids: usr2.GetOrganizationOrUserIdentifiers(),
				Rights: []ttnpb.Right{
					ttnpb.Right_RIGHT_ORGANIZATION_INFO,
					ttnpb.Right_RIGHT_ORGANIZATION_APPLICATIONS_LIST,
					ttnpb.Right_RIGHT_ORGANIZATION_DELETE,
				},
			},
		}, limitedCreds)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		// SetCollaborator removing rights that caller doesn't have.
		_, err = reg.SetCollaborator(ctx, &ttnpb.SetOrganizationCollaboratorRequest{
			OrganizationIds: org1.GetIds(),
			Collaborator: &ttnpb.Collaborator{
				Ids: usr2.GetOrganizationOrUserIdentifiers(),
				Rights: []ttnpb.Right{
					ttnpb.Right_RIGHT_ORGANIZATION_INFO,
				},
			},
		}, limitedCreds)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		// SetCollaborator removing rights that caller has and adding rights that caller has.
		_, err = reg.SetCollaborator(ctx, &ttnpb.SetOrganizationCollaboratorRequest{
			OrganizationIds: org1.GetIds(),
			Collaborator: &ttnpb.Collaborator{
				Ids: usr2.GetOrganizationOrUserIdentifiers(),
				Rights: []ttnpb.Right{
					ttnpb.Right_RIGHT_ORGANIZATION_SETTINGS_BASIC,
					ttnpb.Right_RIGHT_ORGANIZATION_APPLICATIONS_LIST,
				},
			},
		}, limitedCreds)
		a.So(err, should.BeNil)

		// Collaborator CRUD with different invalid credentials.
		for _, opts := range [][]grpc.CallOption{nil, {orgCreds}} {
			_, err := reg.SetCollaborator(ctx, &ttnpb.SetOrganizationCollaboratorRequest{
				OrganizationIds: org1.GetIds(),
				Collaborator: &ttnpb.Collaborator{
					Ids: usr2.GetOrganizationOrUserIdentifiers(),
					Rights: []ttnpb.Right{
						ttnpb.Right_RIGHT_ORGANIZATION_INFO,
					},
				},
			}, opts...)
			if a.So(err, should.NotBeNil) {
				a.So(errors.IsPermissionDenied(err), should.BeTrue)
			}

			got, err := reg.GetCollaborator(ctx, &ttnpb.GetOrganizationCollaboratorRequest{
				OrganizationIds: org1.GetIds(),
				Collaborator:    usr2.GetOrganizationOrUserIdentifiers(),
			}, opts...)
			if a.So(err, should.NotBeNil) && a.So(errors.IsPermissionDenied(err), should.BeTrue) {
				a.So(got, should.BeNil)
			}
		}

		// ListCollaborators without credentials.
		list, err := reg.ListCollaborators(ctx, &ttnpb.ListOrganizationCollaboratorsRequest{
			OrganizationIds: org1.GetIds(),
		})
		if a.So(err, should.NotBeNil) && a.So(errors.IsUnauthenticated(err), should.BeTrue) {
			a.So(list, should.BeNil)
		}

		// Collaborator CRUD with different valid credentials.
		for _, opts := range [][]grpc.CallOption{{adminCreds}, {usr1Creds}, {limitedCreds}} {
			_, err := reg.SetCollaborator(ctx, &ttnpb.SetOrganizationCollaboratorRequest{
				OrganizationIds: org1.GetIds(),
				Collaborator: &ttnpb.Collaborator{
					Ids: usr3.GetOrganizationOrUserIdentifiers(),
					Rights: []ttnpb.Right{
						ttnpb.Right_RIGHT_ORGANIZATION_INFO,
					},
				},
			}, opts...)
			a.So(err, should.BeNil)

			list, err := reg.ListCollaborators(ctx, &ttnpb.ListOrganizationCollaboratorsRequest{
				OrganizationIds: org1.GetIds(),
			}, opts...)
			if a.So(err, should.BeNil) && a.So(list, should.NotBeNil) && a.So(list.Collaborators, should.HaveLength, 3) {
				for _, k := range list.Collaborators {
					if unique.ID(ctx, k.GetIds()) == unique.ID(ctx, usr3.GetIds()) {
						a.So(k.Rights, should.Resemble, []ttnpb.Right{
							ttnpb.Right_RIGHT_ORGANIZATION_INFO,
						})
					}
				}
			}

			got, err := reg.GetCollaborator(ctx, &ttnpb.GetOrganizationCollaboratorRequest{
				OrganizationIds: org1.GetIds(),
				Collaborator:    usr3.GetOrganizationOrUserIdentifiers(),
			}, opts...)
			if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) {
				a.So(got.Rights, should.Resemble, []ttnpb.Right{
					ttnpb.Right_RIGHT_ORGANIZATION_INFO,
				})
			}

			_, err = reg.SetCollaborator(ctx, &ttnpb.SetOrganizationCollaboratorRequest{
				OrganizationIds: org1.GetIds(),
				Collaborator: &ttnpb.Collaborator{
					Ids: usr3.GetOrganizationOrUserIdentifiers(),
				},
			}, opts...)
			a.So(err, should.BeNil)

			got, err = reg.GetCollaborator(ctx, &ttnpb.GetOrganizationCollaboratorRequest{
				OrganizationIds: org1.GetIds(),
				Collaborator:    usr3.GetOrganizationOrUserIdentifiers(),
			}, opts...)
			if a.So(err, should.NotBeNil) {
				a.So(errors.IsNotFound(err), should.BeTrue)
			}
			a.So(got, should.BeNil)
		}

		// Try removing the only collaborator with _ALL rights.
		_, err = reg.SetCollaborator(ctx, &ttnpb.SetOrganizationCollaboratorRequest{
			OrganizationIds: org1.GetIds(),
			Collaborator: &ttnpb.Collaborator{
				Ids: usr1.GetOrganizationOrUserIdentifiers(),
			},
		}, usr1Creds)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsFailedPrecondition(err), should.BeTrue)
		}
	}, withPrivateTestDatabase(p))
}

func TestOrganizationAccessClusterAuth(t *testing.T) {
	p := &storetest.Population{}
	org1 := p.NewOrganization(nil)

	t.Parallel()
	a, ctx := test.New(t)

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		reg := ttnpb.NewOrganizationAccessClient(cc)

		rights, err := reg.ListRights(ctx, org1.GetIds(), is.WithClusterAuth())
		if a.So(err, should.BeNil) && a.So(rights, should.NotBeNil) {
			a.So(ttnpb.AllClusterRights.Intersect(ttnpb.AllOrganizationRights).Sub(rights).Rights, should.BeEmpty)
		}
	}, withPrivateTestDatabase(p))
}
