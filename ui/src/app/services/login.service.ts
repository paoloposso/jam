import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor() { }

  login(email: string, password: string): void {
  }

  isLoggedIn() : boolean {
    return false;
  }
}
