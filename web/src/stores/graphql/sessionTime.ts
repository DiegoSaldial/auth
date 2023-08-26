/* eslint-disable @typescript-eslint/no-explicit-any */
import jwtDecode from 'jwt-decode';

interface TimeSession {
  hours: string;
  minutes: string;
  seconds: string;
}

export async function mostrarTiempoRestante(store: any, refresh = true) {
  const datauser = store.getDataLogin();
  const token = datauser.token;
  if (token) {
    let timeRemaining = getTimeRemaning(token);
    const totalRemaining = getTimeRemaning(datauser.refreshToken);

    if (timeRemaining < 0 && refresh) {
      timeRemaining = getTimeRemaning(token);
    }

    const { hours: a, minutes: b, seconds: c } = getTimeSession(timeRemaining);
    const { hours, minutes, seconds } = getTimeSession(totalRemaining);
    let tiempo_end = '';
    if (hours.includes('-') && minutes.includes('-') && seconds.includes('-')) {
      tiempo_end = `${a}:${b}:${c} <br/> ${a}:${b}:${c}`;
    } else {
      tiempo_end = `${a}:${b}:${c} <br/> ${hours}:${minutes}:${seconds}`;
    }
    store.setExpira(tiempo_end);
  }
  return datauser.dataname;
}

function getTimeSession(timeRemaining: number): TimeSession {
  const hours = (Math.floor(timeRemaining / 3600) + '').padStart(2, '0');
  const min = Math.floor((timeRemaining % 3600) / 60);
  const minutes = (min + '').padStart(2, '0');
  const seconds = (Math.floor(timeRemaining % 60) + '').padStart(2, '0');
  const sesion: TimeSession = {
    hours: hours,
    minutes: minutes,
    seconds: seconds,
  };
  return sesion;
}

function getTimeRemaning(token: string) {
  const decodedToken: any = jwtDecode(token);
  const currentTime = Date.now() / 1000;
  const expirationTime = decodedToken.exp;
  const timeRemaining = expirationTime - currentTime;
  return timeRemaining;
}
