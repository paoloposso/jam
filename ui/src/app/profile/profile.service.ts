import { Injectable } from '@angular/core';
import { Profile } from './profile';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  constructor() { }

  save(profile: Profile): Profile {
    return profile;
  }
}
