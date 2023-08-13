import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RatingComponent } from './rating/rating.component';
import { MatchesComponent } from './matches/matches.component';

const routes: Routes = [
  { path: 'rating', component: RatingComponent },
  { path: 'matches', component: MatchesComponent },
  { path: '',   redirectTo: '/rating', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
