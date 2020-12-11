import { Injectable } from '@angular/core';
import { KO_KR_DICTS } from './ko-kr';

const translateMap = {
  'ko_KR': KO_KR_DICTS,
}

@Injectable()
export class TranslateService {
  private readonly currentLocale = 'ko_KR';

  public translate(key: string): string {
    return translateMap[this.currentLocale][key] || key;
  }  
}