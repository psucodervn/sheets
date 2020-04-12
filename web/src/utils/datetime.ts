import { Month } from '@/types/datetime';

const toMonth = (d: Date): Month => {
  return {
    year: d.getUTCFullYear(),
    month: d.getUTCMonth() + 1,
  };
};

export {
  toMonth,
};
