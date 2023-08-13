import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class EloService {

  constructor() { }

  public async getLeagueName(): Promise<string> {
    return Promise.resolve("TODO")
  }
}
