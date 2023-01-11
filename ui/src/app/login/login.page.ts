import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LoginService } from './login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.page.html',
  styleUrls: ['./login.page.scss'],
})
export class LoginPage implements OnInit {
  user: string;
  pass: string;

  constructor(private service: LoginService, private router: Router) {}

  ngOnInit() {}

  loginUser(){
    this.service.login(this.user, this.pass);
    this.user = "";
    this.pass = "";
  }

  register(){
    this.router.navigate(['register']);
  }
}
