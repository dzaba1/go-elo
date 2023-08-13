import { Component, OnInit } from '@angular/core';
import { EloService } from 'src/services/elo.service';

@Component({
  selector: 'app-matches',
  templateUrl: './matches.component.html',
  styleUrls: ['./matches.component.sass']
})
export class MatchesComponent implements OnInit {
  public displayedColumns = [];
  public isLoaded = false;

  constructor(private eloService: EloService) {

  }

  public async ngOnInit(): Promise<void> {
    this.isLoaded = true;
  }
}
