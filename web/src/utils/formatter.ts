import { Month } from '@/types/datetime';
import { date } from 'quasar';

const formatCurrency = (val: number): string => {
  return Number(val.toFixed(0)).toLocaleString('vi-VN');
};

const formatPoint = (val: number): string => {
  return Number(val.toFixed(1)).toLocaleString();
};

const formatMonth = (d: Month | Date): string => {
  const fm = 'MMMM YYYY';
  if (d instanceof Date) {
    return date.formatDate(d, fm);
  }
  const t = date.buildDate({ year: d.year, month: d.month }, true);
  return date.formatDate(t, fm);
};

export {
  formatCurrency,
  formatPoint,
  formatMonth,
};
