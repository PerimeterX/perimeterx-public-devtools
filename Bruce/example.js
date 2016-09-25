"use strict";
/**
 * Example used in the Bruce blog post [BLOGPOST Link Placeholder]
 *
 * @date October 2016
 * @author ben diamant
 */
const bruceClient = require('./bruce_client').Bruce;
const unix = require('unix-dgram');

const BRUCE_SOCKET = '/opt/bruce/bruce.socket';
const client = unix.createSocket('unix_dgram');

client.on('connect', () => {
    const msg = 'hello from Bruce!';
    const topic = 'bruce_topic';
    const bruceMsg = bruceClient.createAnyPartitionMsg(topic, Date.now(), topic, msg);
    client.send(msg);
});

client.connect(BRUCE_SOCKET);

