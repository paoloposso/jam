import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { ProfileService } from 'src/app/services/profile.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.page.html',
  styleUrls: ['./register.page.scss'],
})
export class RegisterPage implements OnInit {
  registerForm: FormGroup;
  isSubmitted = false;

  email: string;
  name: string;
  password: string;
  passwordConfirmation: string;

  constructor(private formBuilder: FormBuilder, private service: ProfileService) { }

  ngOnInit() {
    this.registerForm = this.formBuilder.group({
      name: new FormControl('', Validators.required),
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', Validators.required),
      passwordConfirmation: new FormControl('', Validators.required),
    });
  }

  onSubmit() {
    this.isSubmitted = true;
    
    if (this.registerForm.valid) {
      this.email = this.registerForm.get('email').value;
      this.name = this.registerForm.get('name').value;
      this.password = this.registerForm.get('password').value;
      this.passwordConfirmation = this.registerForm.get('passwordConfirmation').value;

      this.service.saveBasicData(this.email, this.name, this.password).subscribe(
        p => console.log(p),
        err => alert(err)
      );
    }
  }

  get errorControl() {
    return this.registerForm.controls;
  }
}
