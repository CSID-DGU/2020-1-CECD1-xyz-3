import {
  HttpClient,
  HttpErrorResponse,
} from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'environments/environment';
import * as Rx from 'rxjs';
import * as RxOp from 'rxjs/operators';

type ApplicantResponse = Record<string, Applicant[]>;
type PositionResponse = Record<string, JobPosition[]>;

const TOKEN_STORAGE_KEY = 'DOOGEUN_TOKEN';

interface TokenStorage {
  token: string;
  domain: string;  
}

@Injectable()
export class ApiService {
  private authenticatedDomain: string;
  private commonHeaders: Record<string, string> = {};
  private recentApplicants: ApplicantResponse;
  private recentPositions: PositionResponse;

  public constructor(
    private readonly httpClient: HttpClient,
  ) {
    if (window.localStorage) {
      const tokenStorageString = window.localStorage.getItem(TOKEN_STORAGE_KEY);
      if (tokenStorageString) {
        try {
          const storage: TokenStorage = JSON.parse(tokenStorageString);
          this.commonHeaders['Authorization'] = storage.token;
          this.authenticatedDomain = storage.domain;
        } catch {
        }
      }
    }
  }

  public isAuthenticated(): Rx.Observable<any> {
    return this.httpClient.get(`${environment.apiUrl}/me`, {
      headers: this.commonHeaders,
    });
  }

  public login(userId: string, password: string): Rx.Observable<{ accessToken: string }> {
    return this.httpClient.post<{ accessToken: string }>(`${environment.apiUrl}/token`, {
      email: userId,
    }, {
      observe: 'body',
      responseType: 'json',
    }).pipe(
      RxOp.tap(response => {
        this.authenticatedDomain = userId.substr(userId.lastIndexOf('@') + 1);
        this.commonHeaders['Authorization'] = `Bearer ${response.accessToken}`;
        if (window.localStorage) {
          window.localStorage.setItem(TOKEN_STORAGE_KEY, JSON.stringify({
            token: this.commonHeaders['Authorization'],
            domain: this.authenticatedDomain,
          } as TokenStorage));
        }
      }),
    );
  }

  public logout(): Rx.Observable<any> {
    return new Rx.Observable(observer => {
      this.authenticatedDomain = undefined;
      delete this.commonHeaders['Authorization'];
      if (window.localStorage) {
        window.localStorage.removeItem(TOKEN_STORAGE_KEY);
      }
      observer.next(null);
      observer.complete();
    });
  }

  public getApplicants(companyName?: string): Rx.Observable<Applicant[]> {
    return this.httpClient.get<ApplicantResponse>(`${environment.apiUrl}/applicants`, {
      headers: this.commonHeaders,
      observe: 'body',
      responseType: 'json',
    }).pipe(
      RxOp.catchError((response: HttpErrorResponse) => {
        if (response.status === 404) {
          return Rx.of({});
        }
        return Rx.throwError(response);
      }),
      RxOp.tap(response => this.recentApplicants = response),
      RxOp.map(response => response[companyName || this.authenticatedDomain] || []),
    );
  }

  public getPositions(companyName?: string): Rx.Observable<JobPosition[]> {
    return this.httpClient.get<PositionResponse>(`${environment.apiUrl}/positions`, {
      headers: this.commonHeaders,
      observe: 'body',
      responseType: 'json',
    }).pipe(
      RxOp.catchError((response: HttpErrorResponse) => {
        if (response.status === 404) {
          return Rx.of({});
        }
        return Rx.throwError(response);
      }),
      RxOp.tap(response => this.recentPositions = response),
      RxOp.map(response => response[companyName || this.authenticatedDomain] || []),
    );
  }

  public appendApplicants(applicant: Applicant, fullCompanyName: string): Rx.Observable<any> {
    return new Rx.Observable(observer => {
      this.getApplicants(fullCompanyName)
        .subscribe(response => {
          applicant.id = Math.max(...[ 0, ...response.map(a => a.id) ]) + 1;
          response.push(applicant);
          this.recentApplicants[fullCompanyName] = response;
          this.updateApplicants(this.recentApplicants[fullCompanyName])
            .subscribe(updateResponse => {
              observer.next(null);
              observer.complete();
            }, error => {
              observer.error(error);
            });
        }, error => {
          observer.error(error);
        });
    });
  }

  public updateApplicants(applicants: Applicant[]): Rx.Observable<any> {
    const currentDomainApplicants = this.recentApplicants[this.authenticatedDomain] || [];
    applicants.forEach(applicant => {
      const index = currentDomainApplicants.findIndex(a => a.id === applicant.id);
      if (index < 0) {
        currentDomainApplicants.push(applicant);
      } else {
        currentDomainApplicants[index] = applicant;
      }
    });
    this.recentApplicants[this.authenticatedDomain] = currentDomainApplicants;
    return this.httpClient.put(`${environment.apiUrl}/applicants`, this.recentApplicants, {
      headers: this.commonHeaders,
      observe: 'body',
      responseType: 'text',
    });
  }

  public createPositions(positions: JobPosition[]): Rx.Observable<any> {
    const currentDomainPositions = this.recentPositions[this.authenticatedDomain] || [];
    currentDomainPositions.push(...positions);
    this.recentPositions[this.authenticatedDomain] = currentDomainPositions;
    return this.httpClient.put(`${environment.apiUrl}/positions`, this.recentPositions, {
      headers: this.commonHeaders,
      observe: 'body',
      responseType: 'text',
    });
  }

  public updatePositions(positions: JobPosition[]): Rx.Observable<any> {
    const currentDomainPositions = this.recentPositions[this.authenticatedDomain] || [];
    positions.forEach(position => {
      const index = currentDomainPositions.findIndex(p => p.id === position.id);
      if (index < 0) {
        currentDomainPositions.push(position);
      } else {
        currentDomainPositions[index] = position;
      }
    });
    this.recentPositions[this.authenticatedDomain] = currentDomainPositions;
    return this.httpClient.put(`${environment.apiUrl}/positions`, this.recentPositions, {
      headers: this.commonHeaders,
      observe: 'body',
      responseType: 'text',
    });
  }

  public removePositions(positions: JobPosition[]): Rx.Observable<any> {
    const currentDomainPositions = this.recentPositions[this.authenticatedDomain] || [];
    positions.forEach(position => {
      const index = currentDomainPositions.findIndex(p => p.id === position.id);
      if (index > -1) {
        currentDomainPositions.splice(index, 1);
      }
    });
    this.recentPositions[this.authenticatedDomain] = currentDomainPositions;
    return this.httpClient.put(`${environment.apiUrl}/positions`, this.recentPositions, {
      headers: this.commonHeaders,
      observe: 'body',
      responseType: 'text',
    });
  }
}
