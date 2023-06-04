import { Injectable } from '@angular/core';

export interface Menu {
  state: string;
  name: string;
  type: string;
  icon: string;
}

const MENUITEMS = [
  { state: 'mainpage', type: 'link', name: 'Главная страница', icon: 'menu' },

  { state: 'lists', type: 'link', name: 'Мои Курсы', icon: 'view_list' },
  { state: 'menu', type: 'link', name: 'Курсы', icon: 'view_headline' },
  { state: 'dashboard', name: 'Отзывы', type: 'link', icon: 'av_timer' },
  { state: 'tabs', type: 'link', name: 'О приложении', icon: 'tab' },
];

@Injectable()
export class MenuItems {
  getMenuitem(): Menu[] {
    return MENUITEMS;
  }
}
