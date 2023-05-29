import { Component, OnInit } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-otpiska',
  templateUrl: './otpiska.component.html',
  styleUrls: ['./otpiska.component.css']
})
export class OtpiskaComponent implements OnInit {

  constructor(public dialogRef: MatDialogRef<OtpiskaComponent>) { }

  ngOnInit(): void {
  }

}
