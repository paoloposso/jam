import { ProfileEntity } from "./entities/profile.entity";
import { ProfileRepository } from "./mongo/profile.repository";
import { Inject, Injectable } from "@nestjs/common";
import * as moment from "moment";
import { LocationEntity } from "./entities/location.entity";
import { IProfileRepository } from "./profile.repository";

@Injectable()
export class ProfileService {

    constructor(
        @Inject('IProfileRepository')
        private repository: IProfileRepository) {}
    
    public async getProfileByEmail(email: string): Promise<ProfileEntity> {
        return this.repository.getByEmail(email);
    }

    public async getProfileById(id: string): Promise<ProfileEntity> {
        return this.repository.getById(id);
    }

    public async create(profile: ProfileEntity): Promise<string> {
        let errors: string[] = [];

        this.validateCreateProfileInput(profile, errors);

        profile.registered = this.getUtcDate();

        return this.repository.create(profile);
    }

    public async addInstruments(profileId: string, instruments: string[]) {
        return this.repository.addInstruments(profileId, instruments);
    }

    public async addStyles(profileId: string, styles: string[]) {
        return this.repository.addStyles(profileId, styles);
    }

    public getUtcDate() : Date {
        return moment.utc().toDate();
    }

    public async editAddress(profileId: string, location: LocationEntity) {

    }

    private validateCreateProfileInput(profile: ProfileEntity, errors: string[]) {
        if (!profile.email || profile.email.length === 0) {
            errors.push('Email is required');
        }
        if (!profile.name || profile.name.length === 0) {
            errors.push('Name is required');
        }
        if (errors.length > 0) {
            throw new Error(errors.join(';'));
        }
    }
}