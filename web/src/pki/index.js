import { pki, md } from 'node-forge';

export const getFingerprintSHA1 = certificateData => getFingerprint(certificateData, md.sha1);
export const getFingerprintSHA256 = certificateData => getFingerprint(certificateData, md.sha256);

const getFingerprint = (certificateData, algorithm) => {
    if (!certificateData) {
        return '';
    }

    return pki.getPublicKeyFingerprint(certificateData.publicKey, {
        md: algorithm.create(),
        encoding: 'hex',
        delimiter: ' ',
    }).toUpperCase();
}
