const formatCurrency = (val: number): string => {
  return Number(val.toFixed(0)).toLocaleString('vi-VN');
};

const formatPoint = (val: number): string => {
  return Number(val.toFixed(1)).toLocaleString();
};

export {
  formatCurrency,
  formatPoint,
};
