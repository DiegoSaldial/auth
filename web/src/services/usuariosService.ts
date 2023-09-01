/* eslint-disable @typescript-eslint/no-explicit-any */
import { pageStore } from '../stores/pageStore';
import gql from 'graphql-tag';
import { NewUsuario, UpdateUsuario } from '../types/usuarios';

export default class UsuarioService {
  async usuarios() {
    const store = pageStore();
    const sql = gql`
      query {
        usuarios {
          id
          nombres
          apellidos
          username
          foto_url
          telefono
          correo
          registrado
          estado
          dataname
          documento
          fecha_nac
          domicilio
          roles {
            id
            nombre
            bit
          }
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
  async usuariosByRol(rol_id: string) {
    const store = pageStore();
    const sql = gql`
      query usuariosByRol($rol_id: ID!) {
        usuariosByRol(rol_id: $rol_id) {
          id
          nombres
          apellidos
          username
          foto_url
          telefono
          correo
          registrado
          estado
          dataname
          documento
          fecha_nac
          domicilio
          roles {
            id
            nombre
            bit
          }
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

  async usuarioByUsername(username: string) {
    const store = pageStore();
    const sql = gql`
      query usuarioByUsername($username: String!) {
        usuarioByUsername(username: $username) {
          id
          nombres
          apellidos
          username
          foto_url
          telefono
          correo
          documento
          fecha_nac
          domicilio
          registrado
          estado
          roles {
            id
            nombre
            bit
          }
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { username: username },
      { showNotificacion: true, showNotificacionError: true }
    );
    return res;
  }

  async createUsuario(input: NewUsuario) {
    const store = pageStore();
    const sql = gql`
      mutation createUsuario($input: NewUsuario!) {
        createUsuario(input: $input) {
          id
          nombres
          username
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

  async updateUsuario(input: UpdateUsuario) {
    const store = pageStore();
    const sql = gql`
      mutation updateUsuario($input: UpdateUsuario!) {
        updateUsuario(input: $input) {
          id
          nombres
          username
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
}
