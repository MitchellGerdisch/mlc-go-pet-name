// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Pet
{
    [PetResourceType("pet:index:RandomName")]
    public partial class RandomName : Pulumi.ComponentResource
    {
        /// <summary>
        /// The generated pet name.
        /// </summary>
        [Output("petName")]
        public Output<string> PetName { get; private set; } = null!;


        /// <summary>
        /// Create a RandomName resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RandomName(string name, RandomNameArgs args, ComponentResourceOptions? options = null)
            : base("pet:index:RandomName", name, args ?? new RandomNameArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }

        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class RandomNameArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of parts to pet name.
        /// </summary>
        [Input("numParts", required: true)]
        public Input<double> NumParts { get; set; } = null!;

        public RandomNameArgs()
        {
        }
    }
}