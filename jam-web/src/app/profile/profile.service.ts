import { Injectable } from '@angular/core';
import { Profile } from './profile.model';
import { ApolloModule, Apollo } from 'apollo-angular';
import { HttpLinkModule, HttpLink } from 'apollo-angular-link-http';
import { InMemoryCache } from 'apollo-cache-inmemory';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  constructor(private apollo: Apollo, httpLink: HttpLink) { 
    apollo.create({
      link: httpLink.create({ uri: 'http://localhost:8080' }),
      cache: new InMemoryCache()
    });
  }

  saveProfile(profile: Profile) {
    this.apollo.mutate()
  }
}
