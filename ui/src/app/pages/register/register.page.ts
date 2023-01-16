import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-register',
  templateUrl: './register.page.html',
  styleUrls: ['./register.page.scss'],
})
export class RegisterPage implements OnInit {

  // registerForm: FormGroup;

  constructor() { }

  ngOnInit() {
    // this.registerForm = this.formBuilder.group({
    //   name: new FormControl('', Validators.required),
    //   email: new FormControl('', Validators.required),
    //   gender: new FormControl('', Validators.required),
    //   guitar: new FormControl(false),
    //   bass: new FormControl(false),
    //   drums: new FormControl(false),
    //   keyboard: new FormControl(false),
    // });
  }

}
