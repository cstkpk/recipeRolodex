import * as React from 'react';
import {useDispatch} from 'react-redux';
import {Field, withTypes} from 'react-final-form';

import Select from '../../../components/Select';
import TextInput from '../../../components/TextInput';

import {SearchRecipesForm, validateSearchRecipesForm} from '../formValidation';
import * as actions from '../actions';

import styles from './SearchRecipesForm.scss';

const SearchRecipesForm: React.FC<{}> = () => {
  
  const {Form} = withTypes<SearchRecipesForm>();
  const dispatch = useDispatch();

  const onSubmit = (values: SearchRecipesForm) => {
    const searchQuery: RecipeRolodex.SearchRecipesQuery = {
      ingredient1: values.ingredient1 ? values.ingredient1.trim() : "",
      ingredient2: values.ingredient2 ? values.ingredient2.trim() : "",
      ingredient3: values.ingredient3 ? values.ingredient3.trim() : "",
      season: values.season
    }

    dispatch(actions.getRecipes(searchQuery));
  };

  return (
    <div className={styles.formContainer}>
      <Form onSubmit={onSubmit} validate={validateSearchRecipesForm}>
        {({handleSubmit}) => (
          <form id="form" name="form" onSubmit={handleSubmit} className={styles.form}>
            <Field name="season">
              {({input}) => (
              <Select {...input} label="Season" className={styles.select}>
                <option value="Any">Any Season</option>
                <option value="Winter">Winter</option>
                <option value="Spring">Spring</option>
                <option value="Summer">Summer</option>
                <option value="Fall">Fall</option>
              </Select>
              )}
            </Field>
            <Field name="ingredient1">
              {({input, meta}) => (
                <TextInput 
                  {...input} 
                  label="Ingredient 1"
                  error={meta.error && meta.touched ? meta.error : ''}
                  className={styles.ing1}
                />
              )}
            </Field>
            <Field name="ingredient2">
              {({input, meta}) => (
                <TextInput 
                  {...input} 
                  label="Ingredient 2"
                  error={meta.error && meta.touched ? meta.error : ''}
                  className={styles.ing2}
                />
              )}
            </Field>
            <Field name="ingredient3">
              {({input, meta}) => (
                <TextInput 
                  {...input} 
                  label="Ingredient 3"
                  error={meta.error && meta.touched ? meta.error : ''}
                  className={styles.ing3}
                />
              )}
            </Field>
            <button>Submit</button>
          </form>
        )}
      </Form>
    </div>
  );
};

export default SearchRecipesForm;