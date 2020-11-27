import {PreEnrollmentCode, PreEnrollmentDetails} from "./EnterPreEnrollment";
import {AadharNumber, AadharOTP} from "./AadharVerification";
import React, {createContext, useContext, useMemo, useReducer} from "react";
import {useHistory} from "react-router";

export const FORM_PRE_ENROLL_CODE = "preEnrollCode"
export const FORM_PRE_ENROLL_DETAILS = "preEnrollDetails"
export const FORM_AADHAR_NUMBER = "verifyAadharNumber"
export const FORM_AADHAR_OTP = "verifyAadharOTP"

export function PreEnrollmentFlow(props) {

    function buildForm(pageName) {
        switch (pageName) {
            case FORM_PRE_ENROLL_CODE :
                return <PreEnrollmentCode/>
            case FORM_PRE_ENROLL_DETAILS :
                return <PreEnrollmentDetails/>
            case FORM_AADHAR_NUMBER :
                return <AadharNumber/>
            case FORM_AADHAR_OTP :
                return <AadharOTP/>
        }
        return <PreEnrollmentCode/>
    }

    return (
        <PreEnrollmentProvider>
            {buildForm(props.match.params.pageName)}
        </PreEnrollmentProvider>
    );
}


const PreEnrollmentContext = createContext(null)

export function PreEnrollmentProvider(props) {
    const [state, dispatch] = useReducer(preEnrollmentReducer, initialState)
    const value = useMemo(() => [state, dispatch], [state])
    return <PreEnrollmentContext.Provider value={value} {...props} />
}

const initialState = {};

function preEnrollmentReducer(state, action) {
    switch (action.type) {
        case FORM_PRE_ENROLL_CODE: {
            const newState = {...state}
            newState.enrollCode = action.payload.enrollCode;
            newState.mobileNumber = action.payload.mobileNumber;
            newState.previousForm = action.payload.currentForm;
            return newState
        }
        case FORM_PRE_ENROLL_DETAILS: {
            const newState = {...state}
            newState.previousForm = action.payload.currentForm ?? null
            return newState

        }
        case FORM_AADHAR_NUMBER: {
            const newState = {...state}
            newState.aadharNumber = action.payload.aadharNumber;
            newState.previousForm = action.payload.currentForm ?? null;
            return state
        }
        case FORM_AADHAR_OTP: {
            const newState = {...state}
            newState.aadharOtp = action.payload.aadharOtp;
            newState.previousForm = action.payload.currentForm ?? null;
            return newState
        }
        default:
            throw new Error();
    }
}

export function usePreEnrollment() {

    const context = useContext(PreEnrollmentContext)
    const history = useHistory();
    if (!context) {
        throw new Error(`usePreEnrollment must be used within a PreEnrollmentProvider`)
    }
    const [state, dispatch] = context;

    const goNext = function (current, next, payload) {
        payload.currentForm = current;
        console.log("Next: " + JSON.stringify(payload));
        dispatch({type: current, payload: payload})
        if (next) {
            history.push('/preEnroll/' + next)
        }
    }

    const goBack = function () {
        history.goBack()
    }

    return {
        state,
        dispatch,
        goNext,
        goBack,
    }
}


