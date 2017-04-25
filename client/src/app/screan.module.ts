import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { RouterModule, Routes } from '@angular/router';

import { StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
import { RouterStoreModule } from '@ngrx/router-store';
import { StoreDevtoolsModule } from '@ngrx/store-devtools';

import { SprintService } from './services/sprint.service';

import { SprintEffects } from './store/effects/sprint.effect';

import { reducer } from './store/reducers';

// Pages
import { ScreanComponent } from './screan/screan.component';
import { HomeComponent } from './pages/home/home.component';
import { AddASprintComponent } from './pages/add-a-sprint/add-a-sprint.component';
import { NotFoundComponent } from './pages/not-found/not-found.component';

// Components
import { SprintBoxComponent } from './components/sprint-box/sprint-box.component';

/**
 * Application routes
 */
const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  { path: 'add-a-sprint', component: AddASprintComponent },
  { path: '**', component: NotFoundComponent }
];

@NgModule({
  declarations: [
    HomeComponent,
    SprintBoxComponent,
    AddASprintComponent,
    ScreanComponent,
    NotFoundComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    
    /**
     * Routing
     */
    RouterModule.forRoot(routes),

    /**
     * StoreModule.provideStore is imported once in the root module, accepting a reducer
     * function or object map of reducer functions. If passed an object of
     * reducers, combineReducers will be run creating your application
     * meta-reducer. This returns all providers for an @ngrx/store
     * based application.
     */
    StoreModule.provideStore(reducer),

    /**
     * Store devtools instrument the store retaining past versions of state
     * and recalculating new states. This enables powerful time-travel
     * debugging.
     * 
     * To use the debugger, install the Redux Devtools extension for either
     * Chrome or Firefox
     * 
     * See: https://github.com/zalmoxisus/redux-devtools-extension
     */
    StoreDevtoolsModule.instrumentOnlyWithExtension(),

    /**
     * EffectsModule.run() sets up the effects class to be initialized
     * immediately when the application starts.
     *
     * See: https://github.com/ngrx/effects/blob/master/docs/api.md#run
     */
    EffectsModule.run(SprintEffects)
  ],
  providers: [SprintService],
  bootstrap: [ScreanComponent]
})
export class AppModule { }
