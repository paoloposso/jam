import { Float, Int } from '@nestjs/graphql';
import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document, Types } from 'mongoose';

export type LocationDocument = ProfileLocationModel & Document;

@Schema()
export class ProfileLocationModel {
    constructor(init?: Partial<Location>) {
        Object.assign(this, init);
    }

    @Prop({ type: String, enum: ['Point'], required: false })
    type: string

    @Prop({ type: [Number], required: false })
    coordinates: number[]

    @Prop({ type: String, required: false })
    fullAddress: string
}

export const LocationSchema = SchemaFactory.createForClass(ProfileLocationModel);