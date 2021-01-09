import React from 'react';

import styles from './index.scss';

interface Props {
  onChange: (event: React.ChangeEvent<HTMLInputElement> | any) => void;
  className?: string
  error?: string;
  id?: string;
  label?: string;
  name: string;
  type?: string;
  value: string;
}

const TextInput: React.FC<Props> = (props) => (
  <div className={`${styles.container} ${props.className}`}>
    <label htmlFor={props.id} className={styles.label}>
      {props.label}
    </label>
    <input
      {...props}
      value={props.value || ''}
      className={styles.input}
    />
    <div className={styles.error}>{props.error}</div>
  </div>
);

export default TextInput;