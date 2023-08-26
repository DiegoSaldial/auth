import { Opciones } from '../../types/responses';
import { notificar } from './notificacion';

export function mostrarNotificacion(opciones: Opciones = {}) {
  if (!opciones.showNotificacion) return;
  notificar(
    'Operacion finalizada',
    'La accion ha terminado sin errores',
    'positive'
  );
}
