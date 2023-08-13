import { Component, OnInit } from '@angular/core';
import { Title } from "@angular/platform-browser";
import { EloService } from 'src/services/elo.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent implements OnInit {
  public title: string = ""
  
  constructor(private eloService: EloService,
    private titleService: Title) {
    
  }

  public async ngOnInit(): Promise<void> {
    this.title = await this.eloService.getLeagueName();
    this.titleService.setTitle(this.title);
  }
}
