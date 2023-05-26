import { Component, OnInit } from '@angular/core';
import { BackendService } from '../backend.service';
import { Router } from '@angular/router';
@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {
  username: string = '';
  password: string = '';
  email: string = '';

  signUp(){
    this.back.signUp(this.username,this.password).subscribe((data)=>{
      const token=data.accessToken
      if (token){
        localStorage.setItem('username',this.username)
        localStorage.setItem('token',data.accessToken)
        this.router.navigate(['/dashboard']);
      }
    })
  }

  constructor(private router: Router,private back:BackendService) { }

  ngOnInit(): void {
  }

}
