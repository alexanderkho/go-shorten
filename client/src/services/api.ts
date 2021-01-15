const apiHost = 'http://localhost:8000';

const ping = async (): Promise<string | void> => {
  const r = await fetch(`${apiHost}/api/beep`)
    .then((res) => res.text())
    .catch((e) => console.log(e));
  return r;
};

const postURL = async (url: string): Promise<{id: string}> => {
  const res = await fetch(`${apiHost}/api/url`, {
    method: 'POST',
    body: JSON.stringify({ url }),
  });
  if (res.ok) {
    return res.json();
  }
  const err = await res.text();
  throw new Error(err);
};

export { ping, postURL, apiHost };
