import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OtpiskaComponent } from './otpiska.component';

describe('OtpiskaComponent', () => {
  let component: OtpiskaComponent;
  let fixture: ComponentFixture<OtpiskaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ OtpiskaComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OtpiskaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
