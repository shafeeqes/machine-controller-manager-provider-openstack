/*
Copyright (c) SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMachineSets implements MachineSetInterface
type FakeMachineSets struct {
	Fake *FakeMachineV1alpha1
	ns   string
}

var machinesetsResource = schema.GroupVersionResource{Group: "machine.sapcloud.io", Version: "v1alpha1", Resource: "machinesets"}

var machinesetsKind = schema.GroupVersionKind{Group: "machine.sapcloud.io", Version: "v1alpha1", Kind: "MachineSet"}

// Get takes name of the machineSet, and returns the corresponding machineSet object, and an error if there is any.
func (c *FakeMachineSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MachineSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(machinesetsResource, c.ns, name), &v1alpha1.MachineSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineSet), err
}

// List takes label and field selectors, and returns the list of MachineSets that match those selectors.
func (c *FakeMachineSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MachineSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(machinesetsResource, machinesetsKind, c.ns, opts), &v1alpha1.MachineSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MachineSetList{ListMeta: obj.(*v1alpha1.MachineSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.MachineSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineSets.
func (c *FakeMachineSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(machinesetsResource, c.ns, opts))

}

// Create takes the representation of a machineSet and creates it.  Returns the server's representation of the machineSet, and an error, if there is any.
func (c *FakeMachineSets) Create(ctx context.Context, machineSet *v1alpha1.MachineSet, opts v1.CreateOptions) (result *v1alpha1.MachineSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(machinesetsResource, c.ns, machineSet), &v1alpha1.MachineSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineSet), err
}

// Update takes the representation of a machineSet and updates it. Returns the server's representation of the machineSet, and an error, if there is any.
func (c *FakeMachineSets) Update(ctx context.Context, machineSet *v1alpha1.MachineSet, opts v1.UpdateOptions) (result *v1alpha1.MachineSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(machinesetsResource, c.ns, machineSet), &v1alpha1.MachineSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMachineSets) UpdateStatus(ctx context.Context, machineSet *v1alpha1.MachineSet, opts v1.UpdateOptions) (*v1alpha1.MachineSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(machinesetsResource, "status", c.ns, machineSet), &v1alpha1.MachineSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineSet), err
}

// Delete takes name of the machineSet and deletes it. Returns an error if one occurs.
func (c *FakeMachineSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(machinesetsResource, c.ns, name), &v1alpha1.MachineSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(machinesetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MachineSetList{})
	return err
}

// Patch applies the patch and returns the patched machineSet.
func (c *FakeMachineSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MachineSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(machinesetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MachineSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineSet), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakeMachineSets) UpdateScale(ctx context.Context, machineSetName string, scale *autoscalingv1.Scale, opts v1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(machinesetsResource, "scale", c.ns, scale), &autoscalingv1.Scale{})

	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.Scale), err
}
