// ********RoostGPT********
// Generated Cypress test

// ********RoostGPT********
describe('Voting Application Tests', () => {
    const ecServerEndpoint = 'http://localhost:30081'; // This would be mocked in Cypress

    before(() => {
        cy.intercept('GET', `${ecServerEndpoint}`, { fixture: 'initial-voting-data.json' }).as('getInitialData');
        cy.visit('/voter'); // Assuming the app is served at this path
    });

    it('Successfully loads home page', () => {
        cy.url().should('include', '/voter');
        cy.get('.logo').should('be.visible');
        cy.get('.heading').should('contain', 'Developers preference for building K8S cluster');
    });

    it('Displays initial candidate cards', () => {
        cy.wait('@getInitialData');
        cy.get('.cardContainer').children().should('have.length', 5);
    });

    it('Allows voting and shows notification', () => {
        const randomCandidateIndex = 1; // Simulate voting on second candidate
        cy.get(`.card[data-testid="candidate-${randomCandidateIndex}"]`).click();
        cy.get('.notification').should('contain', 'Vote recorded');
    });

    it('Navigates to results page and shows data', () => {
        cy.get('.results-btn').click();
        cy.url().should('include', '/results');
        cy.get('.resultCard').should('exist');
        cy.get('.totalVotes').should('contain', 'Total Votes:');
    });

    it('Data refreshes automatically with new votes', () => {
        cy.get('.refresh-btn').click();
        cy.wait('@getInitialData'); // Assuming GET call mapped to refresh button
        cy.get('.totalVotes').should('contain', 'Total Votes:');
    });

    it('Handles navigation redirects for unknown routes', () => {
        cy.visit('/voter/unknown');
        cy.url().should('eq', `${Cypress.config().baseUrl}/voter`);
    });

    after(() => {
        cy.clearCookies();
        cy.clearLocalStorage();
    });
});

