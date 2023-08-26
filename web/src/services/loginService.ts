/* eslint-disable @typescript-eslint/no-explicit-any */
import { pageStore } from '../stores/pageStore';
import gql from 'graphql-tag';
import { LoginData } from '../types/login';

export default class LoginService {
  async login(data: LoginData) {
    const store = pageStore();
    const sql = gql`
      mutation login($username: String!, $password: String!) {
        login(username: $username, password: $password) {
          id
          nombre
          bit
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { username: data.username, password: data.password },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }

  async useRol(data: LoginData, rol_id: string) {
    const store = pageStore();
    const sql = gql`
      mutation useRol($username: String!, $password: String!, $rol_id: ID!) {
        useRol(username: $username, password: $password, rol_id: $rol_id) {
          id
          nombres
          apellidos
          dataname
          username
          foto_url
          telefono
          correo
          registrado
          estado
          rol_id
          rol
          exp
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { username: data.username, password: data.password, rol_id: rol_id },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }
}
