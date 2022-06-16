import { LocationEntity } from "./entities/location.entity";
import { ProfileEntity } from "./entities/profile.entity";

export interface IProfileRepository {


    getByEmail(email: string): Promise<ProfileEntity>;

    getById(id: string): Promise<ProfileEntity>;

    create(profile: ProfileEntity): Promise<string>;

    saveLocation(profileId: string, location: LocationEntity): Promise<string>;

    addInstruments(id: string, instruments: string[]);
}