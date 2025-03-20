/**
 * Guarda un valor en el localStorage.
 * @param {string} key - La clave bajo la cual se guardará el valor.
 * @param {any} value - El valor a guardar (será convertido a JSON).
 * @returns {void}
 */
export const saveToLocalStorage = (key: string, value: string) => {
    try {
      const serializedValue = JSON.stringify(value); // Convierte el valor a JSON
      localStorage.setItem(key, serializedValue); // Guarda en localStorage
    } catch (error) {
      console.error("Error al guardar en localStorage:", error);
    }
  };

/**
 * Recupera un valor del localStorage.
 * @param {string} key - La clave del valor a recuperar.
 * @param {any} defaultValue - Valor por defecto si la clave no existe.
 * @returns {any} - El valor recuperado o el valor por defecto.
 */
export const getFromLocalStorage = (key: string, defaultValue = null) => {
    let serializedValue = null;
    try {
    serializedValue = localStorage.getItem(key); // Recupera el valor
      if (serializedValue === null) {
        return defaultValue; // Devuelve el valor por defecto si no existe
      }

      return JSON.parse(serializedValue); // Convierte de JSON a objeto/valor
    } catch (error) {
      console.error("Error al recuperar de localStorage:", error);
      return serializedValue || defaultValue; // Devuelve el valor string O el valor por defecto
    }
  };