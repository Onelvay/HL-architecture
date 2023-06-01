import { Component, OnInit } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { BackendService } from '../backend.service';
@Component({
  selector: 'app-review-dialog',
  templateUrl: './review-dialog.component.html',
  styleUrls: ['./review-dialog.component.css']
})
export class ReviewDialogComponent implements OnInit {
  comments: string="Комментарии";
  rating: number=5;
  reviewForm: FormGroup=new FormGroup({
    comments: new FormControl('', Validators.required),
    rating: new FormControl('', [Validators.required, Validators.min(1), Validators.max(5)])
  });;
  constructor(public dialogRef: MatDialogRef<ReviewDialogComponent>,private back:BackendService) {}

  submitReview(): void {
    if (this.reviewForm.valid) {
      // Здесь вы можете добавить логику для отправки отзыва на сервер или выполнения других необходимых действий
      console.log('Отправлен отзыв:', this.reviewForm.value);
      const token =localStorage.getItem('token')
      const orderId =localStorage.getItem('orderId')
      if (token && orderId){
        console.log(this.reviewForm.value[0],this.reviewForm.value[1])
        this.back.addReview(token,orderId,this.reviewForm.value.comments,this.reviewForm.value.rating).subscribe()
        this.dialogRef.close();
      }
     

    }
  }
  ngOnInit(): void {
    
  }

}
