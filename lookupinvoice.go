package main

import (
	"github.com/bumi/lookupinvoice"
)

func lookupInvoice(
	params *Params,
	paymentHash string,
) (invoice lookupinvoice.Invoice, err error) {
	// prepare params
	var backend lookupinvoice.BackendParams
	switch params.Kind {
	case "sparko":
		backend = lookupinvoice.SparkoParams{
			Host: params.Host,
			Key:  params.Key,
		}
	case "lnd":
		backend = lookupinvoice.LNDParams{
			Host:     params.Host,
			Macaroon: params.Key,
		}
	case "lnbits":
		backend = lookupinvoice.LNBitsParams{
			Host: params.Host,
			Key:  params.Key,
		}
	}

	lip := lookupinvoice.Params{
		Backend:     backend,
		PaymentHash: paymentHash,
	}

	// actually generate the invoice
	invoice, err = lookupinvoice.LookupInvoice(lip)

	log.Debug().Str("paymentHash", paymentHash).
		Interface("backend", backend).
		Interface("invoice", invoice).Err(err).
		Msg("invoice lookup")

	return invoice, err
}
