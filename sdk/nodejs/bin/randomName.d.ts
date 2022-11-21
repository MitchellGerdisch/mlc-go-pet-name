import * as pulumi from "@pulumi/pulumi";
export declare class RandomName extends pulumi.ComponentResource {
    /**
     * Returns true if the given object is an instance of RandomName.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is RandomName;
    /**
     * The generated pet name.
     */
    readonly petName: pulumi.Output<string>;
    /**
     * Create a RandomName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RandomNameArgs, opts?: pulumi.ComponentResourceOptions);
}
/**
 * The set of arguments for constructing a RandomName resource.
 */
export interface RandomNameArgs {
    /**
     * Number of parts to pet name.
     */
    numParts: pulumi.Input<number>;
}
