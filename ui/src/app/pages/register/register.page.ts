import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-register',
  templateUrl: './register.page.html',
  styleUrls: ['./register.page.scss'],
})
export class RegisterPage implements OnInit {

  registerForm: FormGroup;

  constructor(private formBuilder: FormBuilder) { }

  ngOnInit() {
    this.registerForm = this.formBuilder.group({
      name: new FormControl('', Validators.required),
      email: new FormControl('', Validators.required),
      password: new FormControl('', Validators.required),
      passwordConfirm: new FormControl('', Validators.required),
    });
  }

  onSubmit() {
    alert('OK');
    // this.isSubmitted = true;
    // if (this.profileForm.valid) {
    //   let profile = new Profile();

    //   profile.email = this.profileForm.get('email').value;
    //   profile.name = this.profileForm.get('name').value;
    //   profile.gender = this.profileForm.get('gender').value;

    //   this.service.save(profile).subscribe(
    //     p => console.log(p),
    //     err => alert(err)
    //   );
    // }
  }
}
