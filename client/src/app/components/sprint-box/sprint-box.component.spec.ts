import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { SprintBoxComponent } from './sprint-box.component';

describe('SprintBoxComponent', () => {
  let component: SprintBoxComponent;
  let fixture: ComponentFixture<SprintBoxComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SprintBoxComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SprintBoxComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
