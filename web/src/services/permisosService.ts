/* eslint-disable @typescript-eslint/no-explicit-any */
import { pageStore } from '../stores/pageStore';
import gql from 'graphql-tag';

export default class PermisosService {
  async permisosByRol(rol_id: string) {
    const store = pageStore();
    const sql = gql`
      query permisosByRol($rol_id: ID!) {
        permisosByRol(rol_id: $rol_id) {
          metodo
          descripcion
          rol_bits
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { rol_id: rol_id },
      { showNotificacion: true, showNotificacionError: true }
    );
    return res;
  }

  async permisos() {
    const store = pageStore();
    const sql = gql`
      query permisos {
        permisos {
          metodo
          descripcion
          rol_bits
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      {},
      { showNotificacion: true, showNotificacionError: true }
    );
    return res;
  }
}
