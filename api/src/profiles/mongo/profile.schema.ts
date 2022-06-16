import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document, Types } from 'mongoose';
import { ProfileLocationModel } from './location.schema';

export type ProfileDocument = ProfileModel & Document;

@Schema()
export class ProfileModel {
    constructor(init?: Partial<ProfileModel>) {
        Object.assign(this, init);
    }

    @Prop({ type: Types.ObjectId })
    id: string;

    @Prop({required: true, unique: true, index: true})
    email: string;

    @Prop({ required: true })
    name: string;

    @Prop({ index: true, required: false })
    styles: string[];

    @Prop({ index: true, required: false })
    instruments: string[];

    @Prop({ required: true })
    registered: Date;

    @Prop({ required: false })
    location: ProfileLocationModel;    
}

export const ProfileSchema = SchemaFactory.createForClass(ProfileModel);