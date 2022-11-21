// Copyright 2016-2021, Pulumi Corporation.
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

package provider

import (
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The set of arguments for creating a RandomName component resource.
type RandomNameArgs struct {
	// The HTML content for index.html.
	NumParts pulumi.IntPtrInput `pulumi:"numParts"`
}

// The RandomName component resource.
type RandomName struct {
	pulumi.ResourceState

	PetName pulumi.StringOutput `pulumi:"petName"`
}

// NewRandomName creates a new RandomName component resource.
func NewRandomName(ctx *pulumi.Context,
	name string, args *RandomNameArgs, opts ...pulumi.ResourceOption) (*RandomName, error) {
	if args == nil {
		args = &RandomNameArgs{}
	}

	component := &RandomName{}
	err := ctx.RegisterComponentResource("pet:index:RandomName", name, component, opts...)
	if err != nil {
		return nil, err
	}

	// Create a random pet name
	pet, err := random.NewRandomPet(ctx, name, &random.RandomPetArgs{
		Length: args.NumParts,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	component.PetName = pulumi.StringOutput(pulumi.IDOutput(pet.ID()))

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"petName": component.PetName,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
