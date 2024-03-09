import React, {useEffect} from 'react';
import './App.css';
import {Alert, Box, Button, Grid, Paper, TextField, Typography} from "@mui/material";
import {Pack} from "./model/pack";
import {API_URL} from "./config";
import {ShippingResponse} from "./model/shipping-response";

function App() {
    const [packs, setPacks] = React.useState<Pack[]>([]);
    const [shippingResult, setShippingResult] = React.useState<ShippingResponse | null>(null);
    const [itemCount, setItemCount] = React.useState<number | null>(0);
    const [error, setError] = React.useState<string | null>(null);

    useEffect(() => {
        fetch(API_URL + '/packs')
            .then(resp => {
                return resp.json()
            })
            .catch(err => setError(err))
            .then(data => {
                setPacks(data)
            })
    }, [])

    function calculateShipping() {
        fetch(API_URL + '/shipping', {
            method: 'POST',
            body: JSON.stringify({itemCount}),
        })
            .then(resp => resp.json())
            .catch(err => setError(err))
            .then(data => {
                setShippingResult(data)
            })
    }

    return (
        <div className="App">
            <Box maxWidth='600px' margin='40px auto' display='flex'>
                <Paper sx={{
                    width: '100%',
                }}>
                    <Grid width='100%' container spacing={2}>
                        <Grid item xs={2}>
                            <Typography variant='h5'>Packs:</Typography>
                            <ul>
                                {packs.map(pack => {
                                    return <li key={pack.id}>{pack.itemCount}</li>
                                })}
                            </ul>
                        </Grid>
                        <Grid item xs={10}>
                            <Typography variant='h5'>Shipping Pack calculator</Typography>
                            <br/>
                            <TextField value={itemCount} onChange={e => {
                                setItemCount(parseInt(e.target.value))
                            }} type='number' label='Item count' variant='outlined'/>
                            <Button onClick={() => calculateShipping()}>Calculate</Button>
                            <br/>
                            <br/>

                            {!error && shippingResult && (
                                <div>
                                    <Typography variant='h5'>Shipping result</Typography>
                                    {shippingResult.shipping.items && (
                                        <ul>
                                            {shippingResult.shipping.items.map(item => {
                                                return <li key={item.pack.id}>{item.count}x{item.pack.itemCount} </li>
                                            })}
                                        </ul>
                                    )}
                                </div>
                            )}
                            {error && <Alert severity="error">{error}</Alert>}
                        </Grid>
                    </Grid>
                </Paper>
            </Box>
        </div>
    );
}

export default App;
