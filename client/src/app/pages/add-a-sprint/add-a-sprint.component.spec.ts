import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AddASprintComponent } from './add-a-sprint.component';

describe('AddASprintComponent', () => {
  let component: AddASprintComponent;
  let fixture: ComponentFixture<AddASprintComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AddASprintComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AddASprintComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
