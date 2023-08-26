/* eslint-disable @typescript-eslint/no-explicit-any */
import { pageStore } from '../stores/pageStore';
import gql from 'graphql-tag';
import { ModificarRol } from '../types/roles';

export default class RolesService {
  async roles(lite: boolean) {
    const store = pageStore();
    const sql = gql`
      query roles($lite: Boolean!) {
        roles(lite: $lite) {
          id
          nombre
          bit
          permisos
          usuarios
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { lite: lite },
      { showNotificacion: true, showNotificacionError: true }
    );
    return res;
  }

  async createRol(input: ModificarRol) {
    const store = pageStore();
    const sql = gql`
      mutation createRol($input: NewRol!) {
        createRol(input: $input) {
          id
          nombre
          bit
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { input: input },
      { showNotificacion: true, showNotificacionError: true }
    );
    return res;
  }

  async modificarRol(input: ModificarRol) {
    const store = pageStore();
    const sql = gql`
      mutation modificarRol($input: UpdateRol!) {
        modificarRol(input: $input) {
          id
          nombre
          bit
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { input: input },
      { showNotificacion: true, showNotificacionError: true }
    );
    return res;
  }

  async deleteRol(rol_id: string) {
    const store = pageStore();
    const sql = gql`
      mutation deleteRol($rol_id: ID!) {
        deleteRol(rol_id: $rol_id)
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { rol_id: rol_id },
      { showNotificacion: true, showNotificacionError: true }
    );
    return res;
  }
}
