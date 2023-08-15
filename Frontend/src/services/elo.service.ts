import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Match } from '../model/match';
import { PlayerRating } from '../model/playerRating';
import { firstValueFrom } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class EloService {

  private baseUrl = "http://localhost:8080/";

  constructor(private http: HttpClient) {
  }

  public async getLeagueName(): Promise<string> {
    const resp = await firstValueFrom(this.http.get<LeagueNameResp>(this.baseUrl + "name"))
    return resp.name;
  }

  public async getRating(): Promise<PlayerRating[]> {
    return await firstValueFrom(this.http.get<PlayerRating[]>(this.baseUrl + "elo"))
  }

  public async addMatch(match: Match): Promise<Match> {
    return await firstValueFrom(this.http.post<Match>(this.baseUrl + "match", match))
  }

  public async getMatches(): Promise<Match[]> {
    return await firstValueFrom(this.http.get<Match[]>(this.baseUrl + "match"))
  }

  public async deleteMatch(id: number): Promise<void> {
    await firstValueFrom(this.http.delete(this.baseUrl + `match/${id}`));
  }
}

interface LeagueNameResp {
  name: string
}