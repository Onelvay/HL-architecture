import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import {MatListModule} from '@angular/material/list';
import { BackendService } from 'src/app/backend.service';
import { ConfirmComponent } from 'src/app/confirm/confirm.component';
import { Course ,Order} from 'src/app/models';
import { OtpiskaComponent } from 'src/app/otpiska/otpiska.component';

@Component({
  selector: 'app-lists',
  templateUrl: './lists.component.html',
  styleUrls: ['./lists.component.scss']
})
export class ListsComponent {
  courses:Course[]=[];
  constructor(private back:BackendService,private dialog: MatDialog){}
  ngOnInit(){
    const token  = localStorage.getItem('token')
    console.log(token)
    if (token){
      this.back.getUserCourses(token).subscribe((val)=>{
        this.courses=val
      })

    }
  }
  testt(orderId:string):void{
    const dialogRef = this.dialog.open(OtpiskaComponent);

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        const token = localStorage.getItem("token")
        if (token){
          this.back.deleteOrder(token,orderId).subscribe()
        }// Выполните действия при подтверждении регистрации
       
      }
  });
}
}
