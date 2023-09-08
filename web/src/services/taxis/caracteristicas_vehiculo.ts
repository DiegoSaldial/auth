/* eslint-disable @typescript-eslint/no-explicit-any */
import { pageStore } from '../../stores/pageStore';
import gql from 'graphql-tag';
import { CreateCaractericaVehiculos } from '../../types/taxis/caracteristicas_vehiculo';

export default class CaracteristicasVehiculosService {
  async caracteristicas() {
    const store = pageStore();
    const sql = gql`
      query{
        caracteristicas{id nombre registro vehiculos}
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }
  async caracteristica(id:string) {
    const store = pageStore();
    const sql = gql`
      query caracteristica($id:ID!){
        caracteristica(id:$id){id nombre registro vehiculos}
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { id:id},
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }
  async caracteristicasByVehiculo(id:string) {
    const store = pageStore();
    const sql = gql`
      query caracteristicasByVehiculo($id:ID!){
        caracteristicasByVehiculo(id:$id){id nombre registro vehiculos}
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { id:id},
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }

  async createUpdateCaracteristica(input:CreateCaractericaVehiculos) {
    const store = pageStore();
    const sql = gql`
      mutation createUpdateCaracteristica($input:NewCaracteristica!){
        createUpdateCaracteristica(input:$input){id nombre registro vehiculos}
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { input:input },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }


}
