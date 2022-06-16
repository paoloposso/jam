import { Field, ID, Int, ObjectType } from '@nestjs/graphql';

@ObjectType()
export class LocationEntity {

    constructor(init?: Partial<LocationEntity>) {
        Object.assign(this, init);
    }

    @Field()
    type: String

    @Field({ nullable: false })
    coordinates: number[]

    @Field({ nullable: false })
    fullAddress: string
}