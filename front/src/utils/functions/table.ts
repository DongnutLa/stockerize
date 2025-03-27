const pageOptions = [5, 10, 20, 50, 100, 200, 500];

export function getPageSizeOptions(totalData: number): number[] {
  let valorMasCercano = pageOptions[0];

  for (const option of pageOptions) {
    if (option >= totalData) {
      valorMasCercano = option;
      break;
    }
  }

  return pageOptions.filter((option) => option <= valorMasCercano);
}