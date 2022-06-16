import { Field, ID, Int, ObjectType } from '@nestjs/graphql';

@ObjectType()
export class LocationEntity {

    constructor(init?: Partial<LocationEntity>) {
        Object.assign(this, init);
    }

    @Field({ nullable: true })
    type: string;

    @Field(_type => [Number], { nullable: true })
    coordinates: number[] = [];

    @Field({ nullable: true })
    fullAddress: string;
}