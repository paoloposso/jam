import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor() { }

  login(user: string, pass: string): void {
  }

  isLoggedIn() : boolean {
    return false;
  }
}
