import { Month } from "@/types/datetime";
import { date as dateUtils } from "quasar";

const currency = (val: number): string => {
  return Number(val.toFixed(0)).toLocaleString("vi-VN");
};

const storyPoint = (val: number): string => {
  return Number(val.toFixed(1)).toLocaleString();
};

const month = (d: Month | Date): string => {
  const fm = "MMMM YYYY";
  if (d instanceof Date) {
    return dateUtils.formatDate(d, fm);
  }
  const t = dateUtils.buildDate({ year: d.year, month: d.month }, true);
  return dateUtils.formatDate(t, fm);
};

const date = (d: Date): string => {
  return dateUtils.formatDate(d, "DD/MM");
};

const formatter = {
  currency,
  storyPoint,
  month,
  date
};

export default formatter;
