import { Component, OnInit } from '@angular/core';
import {STEPPER_GLOBAL_OPTIONS} from '@angular/cdk/stepper';
import { Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BackendService } from '../backend.service';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  
})
export class LoginComponent implements OnInit {
  logged: boolean = false;
  username: string = '';
  password: string = '';
  right:boolean=true;
  constructor(private router: Router,private back:BackendService) {}


  ngOnInit() {
    const token = localStorage.getItem('token');
    if (token) {
      this.logged = true;
    }

  }

  login() {
    this.back.signIn(this.username, this.password).subscribe((data) => {
      localStorage.setItem('token', data.accessToken);
      localStorage.setItem('username',this.username);
      this.logged = true;
    });
    const token = localStorage.getItem('token');
    if(token){
      this.router.navigate(['/dashboard']);
    }else{
      this.right=false;
    }
  }
}
