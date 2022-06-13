import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document, Types } from 'mongoose';

export type ProfileDocument = Profile & Document;

@Schema()
export class Profile {
    constructor(init?: Partial<Profile>) {
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
}

export const ProfileSchema = SchemaFactory.createForClass(Profile);