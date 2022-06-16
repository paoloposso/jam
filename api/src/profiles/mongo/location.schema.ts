import { Float, Int } from '@nestjs/graphql';
import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document, Types } from 'mongoose';

export type LocationDocument = Location & Document;

@Schema()
export class Location {
    constructor(init?: Partial<Location>) {
        Object.assign(this, init);
    }

    @Prop({ type: String, enum: ['Point'], required: true })
    type: String

    @Prop({ type: [Float], required: true })
    coordinates: number[]

    @Prop({ type: [String], required: true })
    fullAddress: string
}

export const LocationSchema = SchemaFactory.createForClass(Location);